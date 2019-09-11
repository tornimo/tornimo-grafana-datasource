package main

type TimePoint [2]*float64

type TargetResponseDTO struct {
	Target     string      `json:"target"`
	DataPoints []TimePoint `json:"datapoints"`
}
