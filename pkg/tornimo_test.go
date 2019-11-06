/**
 *  Copyright 2015 Grafana Labs
 *  Modifications Copyright 2019 Tornimo
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

 package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTornimoFunctions(t *testing.T) {
	Convey("Testing Tornimo Functions", t, func() {

		Convey("formatting time range for now", func() {

			timeRange := formatTimeRange("now")
			So(timeRange, ShouldEqual, "now")

		})

		Convey("formatting time range for now-1m", func() {

			timeRange := formatTimeRange("now-1m")
			So(timeRange, ShouldEqual, "-1min")

		})

		Convey("formatting time range for now-1M", func() {

			timeRange := formatTimeRange("now-1M")
			So(timeRange, ShouldEqual, "-1mon")

		})

		Convey("fix interval format in query for 1m", func() {

			timeRange := fixIntervalFormat("aliasByNode(hitcount(averageSeries(app.grafana.*.dashboards.views.count), '1m'), 4)")
			So(timeRange, ShouldEqual, "aliasByNode(hitcount(averageSeries(app.grafana.*.dashboards.views.count), '1min'), 4)")

		})

		Convey("fix interval format in query for 1M", func() {

			timeRange := fixIntervalFormat("aliasByNode(hitcount(averageSeries(app.grafana.*.dashboards.views.count), '1M'), 4)")
			So(timeRange, ShouldEqual, "aliasByNode(hitcount(averageSeries(app.grafana.*.dashboards.views.count), '1mon'), 4)")

		})

		Convey("should not override query for 1M", func() {

			timeRange := fixIntervalFormat("app.grafana.*.dashboards.views.1M.count")
			So(timeRange, ShouldEqual, "app.grafana.*.dashboards.views.1M.count")

		})

		Convey("should not override query for 1m", func() {

			timeRange := fixIntervalFormat("app.grafana.*.dashboards.views.1m.count")
			So(timeRange, ShouldEqual, "app.grafana.*.dashboards.views.1m.count")

		})
	})
}
