

type Matrix struct {

	Rows, Cols [][]int 
	Data []float64


}

func (m *Matrix) (at i,j int) float64{

	return m.Data[i*m.Cols + j]

}
