package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"

	"golang.org/x/net/context/ctxhttp"

	"github.com/opentracing/opentracing-go"
	"github.com/hashicorp/go-hclog"
	"github.com/grafana/grafana_plugin_model/go/datasource"
	"github.com/hashicorp/go-plugin"
)

type TornimoDatasource struct {
	plugin.NetRPCUnsupportedPlugin
	HttpClient *http.Client
}

var (
	logger = hclog.New(&hclog.LoggerOptions{
		Name:  "tsdb.tornimo",
		Level: hclog.LevelFromString("DEBUG"),
	})
)

func (e *TornimoDatasource) Query(ctx context.Context, request *datasource.DatasourceRequest) (*datasource.DatasourceResponse, error) {

	from := "-" + formatTimeRange(request.TimeRange.FromRaw)
	until := formatTimeRange(request.TimeRange.ToRaw)
	var target string

	var dataSourceJsonData map[string]interface{}
	json.Unmarshal([]byte(request.Datasource.JsonData), &dataSourceJsonData)
	var token = dataSourceJsonData["token"].(string)

	formData := url.Values{
		"from":          []string{from},
		"until":         []string{until},
		"format":        []string{"json"},
		"maxDataPoints": []string{"500"},
		"token":         []string{token},
	}

	for _, query := range request.Queries {
		var jsonModel map[string]interface{}
		json.Unmarshal([]byte(query.ModelJson), &jsonModel)

		if jsonModel["targetFull"] != nil {
			target = fixIntervalFormat(jsonModel["targetFull"].(string))
		} else {
			target = fixIntervalFormat(jsonModel["target"].(string))
		}
	}

	formData["target"] = []string{target}

	req, err := e.createRequest(request.Datasource, token, formData)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "tornimo query")
	span.SetTag("target", target)
	span.SetTag("from", from)
	span.SetTag("until", until)
	defer span.Finish()

	opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header))

	res, err := ctxhttp.Do(ctx, e.HttpClient, req)
	if err != nil {
		return nil, err
	}

	data, err := e.parseResponse(res)
	if err != nil {
		return nil, err
	}

	var timesSeries []*datasource.TimeSeries
	for _, series := range data {
		logger.Info("Time series", "ts", datasource.TimeSeries{
			Name:   series.Target,
			Points: parsePoints(series.DataPoints),
		})

		timesSeries = append(timesSeries, &datasource.TimeSeries{
			Name:   series.Target,
			Points: parsePoints(series.DataPoints),
		})
	}

	var newFormatQueryResult = &datasource.QueryResult{
		Series: timesSeries,
	}

	return &datasource.DatasourceResponse{
		Results: []*datasource.QueryResult{newFormatQueryResult},
	}, nil
}

func (e *TornimoDatasource) parseResponse(res *http.Response) ([]TargetResponseDTO, error) {
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	if res.StatusCode/100 != 2 {
		logger.Info("Request failed", "status", res.Status, "body", string(body))
		return nil, fmt.Errorf("Request failed status: %v", res.Status)
	}

	var data []TargetResponseDTO
	err = json.Unmarshal(body, &data)
	if err != nil {
		logger.Info("Failed to unmarshal tornimo response", "error", err, "status", res.Status, "body", string(body))
		return nil, err
	}

	return data, nil
}

func (e *TornimoDatasource) createRequest(dsInfo *datasource.DatasourceInfo, token string, data url.Values) (*http.Request, error) {
	u, _ := url.Parse(dsInfo.Url)
	u.Path = path.Join(u.Path, "render")

	req, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(data.Encode()))
	if err != nil {
		logger.Info("Failed to create request", "error", err)
		return nil, fmt.Errorf("Failed to create request. error: %v", err)
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, err
}

func formatTimeRange(input string) string {
	if input == "now" {
		return input
	}
	return strings.Replace(strings.Replace(strings.Replace(input, "now", "", -1), "m", "min", -1), "M", "mon", -1)
}

func fixIntervalFormat(target string) string {
	rMinute := regexp.MustCompile(`'(\d+)m'`)
	rMin := regexp.MustCompile("m")
	target = rMinute.ReplaceAllStringFunc(target, func(m string) string {
		return rMin.ReplaceAllString(m, "min")
	})
	rMonth := regexp.MustCompile(`'(\d+)M'`)
	rMon := regexp.MustCompile("M")
	target = rMonth.ReplaceAllStringFunc(target, func(M string) string {
		return rMon.ReplaceAllString(M, "mon")
	})
	return target
}

func parsePoints(timeSeriesPoints []TimePoint) []*datasource.Point {
	var newPoints []*datasource.Point
	for _, point := range timeSeriesPoints {
		newPoint := parseTimePoint(point)
		newPoints = append(newPoints, newPoint)
	}
	return newPoints
}

func parseTimePoint(point TimePoint) *datasource.Point {
	var newPoint *datasource.Point
	if point[0] != nil {
		newPoint = &datasource.Point{
			Value:     *point[0],
			Timestamp: int64(*point[1]),
		}

	} else {
		newPoint = &datasource.Point{
			Timestamp: int64(*point[1]),
		}
	}
	return newPoint
}
