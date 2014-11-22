package gorilla

import "fmt"
import linuxproc "github.com/c9s/goprocinfo/linux"

type LoadAverage struct {
        last1m, last5m, last15m float64
}

func (l* LoadAverage) Usage() float64 {
        loadAverage, _ := linuxproc.ReadLoadAvg("/proc/loadavg")

        l.last1m = loadAverage.Last1Min
        l.last5m = loadAverage.Last5Min
        l.last15m = loadAverage.Last15Min

        return l.last1m
}

func (l* LoadAverage) Ranking() string {
        return fmt.Sprintf("1-minute load average/core is %f", l.last1m)
}

func (l* LoadAverage) Report(metricQueue chan *Metric){

        metric := new(Metric)

        metric.service = "load"
        metric.value = l.Usage()
        metric.description = l.Ranking()

        metricQueue <- metric
}
