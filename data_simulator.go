package query

import (
	"github.com/golang/protobuf/ptypes"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
	"math"
	"os"
	"time"
)

func sequence(length int) (xs []float64) {
	xs = make([]float64, length, length)
	for i := range xs {
		xs[i] = float64(i) / (float64(length - 1))
	}
	return
}

//Trig spec for trigonometry curve
type Trig struct {
	amplitude float64
	periods   float64
	hOffset   float64
	vOffset   float64
}

func NewTrig(
	amplitude float64,
	period float64,
	hOffset float64,
	vOffset float64) (t Trig) {
	return Trig{
		amplitude: amplitude,
		periods:   period,
		hOffset:   hOffset,
		vOffset:   vOffset,
	}
}

func (t Trig) sinArr(xs *[]float64) {
	for i := range *xs {
		(*xs)[i] = t.amplitude*math.Sin(2.0*math.Pi*t.periods*(t.hOffset+float64(i)/float64(len(*xs)-1))) + t.vOffset
	}
	return
}

func (t Trig) timeSeries(first time.Time, last time.Time) (out []DataPoint, err error) {
	n := int(last.Sub(first).Hours() / 24)
	y := sequence(n)
	t.sinArr(&y)
	out = make([]DataPoint, n, n)
	for i := range y {
		timestamp, err := ptypes.TimestampProto(first.AddDate(0, 0, i))
		if err != nil {
			return nil, err
		}
		out[i] = DataPoint{
			Timestamp: timestamp,
			Value:     y[i],
		}
	}
	return out, nil
}

func (t Trig) timeSeriesPtrs(first time.Time, last time.Time) (out []*DataPoint, err error) {
	n := int(last.Sub(first).Hours() / 24)
	y := sequence(n)
	t.sinArr(&y)
	for i := range y {
		timestamp, err := ptypes.TimestampProto(first.AddDate(0, 0, i))
		if err != nil {
			return nil, err
		}
		out = append(out, &DataPoint{
			Timestamp: timestamp,
			Value:     y[i],
		})
	}
	return out, nil
}

func (t Trig) timeSeriesPlotter(n int) plotter.XYs {
	y := sequence(n)
	t.sinArr(&y)
	const (
		year  = 2017
		month = 1
		day   = 1
		hour  = 1
		min   = 1
		sec   = 1
		nsec  = 1
	)
	pts := make(plotter.XYs, n)
	for i := range pts {
		date := time.Date(year, month, day, hour, min, sec, nsec, time.UTC).Add(time.Hour * time.Duration(24*i)).Unix()
		pts[i].X = float64(date)
		pts[i].Y = y[i]
	}
	return pts
}

func plotFoo() {
	n := 100
	t := NewTrig(5, 1, 0, 10)
	name := "FOO"
	linePointsData := t.timeSeriesPlotter(n)

	// Create a new plot, set its title and axis labels.
	p := plot.New()
	p.Title.Text = "Asset Price Time Series"
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Price"
	p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02"}
	// Draw a grid behind the data
	p.Add(plotter.NewGrid())

	// Make a line plotter with points and set its style.
	lpLine, lpPoints, err := plotter.NewLinePoints(linePointsData)
	if err != nil {
		panic(err)
	}
	lpLine.Color = color.RGBA{G: 255, A: 255}
	lpPoints.Shape = draw.PyramidGlyph{}
	lpPoints.Color = color.RGBA{R: 255, A: 255}

	// Add the plotters to the plot, with a legend
	// entry for each
	p.Add(lpLine, lpPoints)
	p.Legend.Add(name, lpLine, lpPoints)

	outpath := "points.png"
	// Save the plot to a PNG file.
	os.Remove(outpath)
	if err := p.Save(10*vg.Centimeter, 10*vg.Centimeter, outpath); err != nil {
		panic(err)
	}
}
