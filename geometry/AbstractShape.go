package geometry

type AbstractShape struct {
	id       string
	center   *Vector2
	radius   float64
	userData interface{}
}

func (a *AbstractShape) GetID() string {
	return a.id
}

func (a *AbstractShape) GetCenter() *Vector2 {
	return a.center
}

func (a *AbstractShape) GetRadius() float64 {
	return a.radius
}

func (a *AbstractShape) GetUserData() interface{} {
	return a.userData
}

func (a *AbstractShape) SetUserData(data interface{}) {
	a.userData = data
}

func (a *AbstractShape) RotateAboutXY(theta, x, y float64) {
	if !(a.center.X == x && a.center.Y == y) {
		a.center.RotateAboutXY(theta, x, y)
	}
}

func (a *AbstractShape) TranslateXY(x, y float64) {
	a.center.AddXY(x, y)
}

/*
func (a *AbstractShape) RotateAboutOrigin(theta float64) {
	a.RotateAboutXY(theta, 0, 0)
}

func (a *AbstractShape) RotateAboutCenter(theta float64) {
	a.RotateAboutXY(theta, a.center.X, a.center.Y)
}

func (a *AbstractShape) RotateAboutVector2(theta float64, v *Vector2) {
	a.RotateAboutXY(theta, v.X, v.Y)
}

func (a *AbstractShape) TranslateVector2(v *Vector2) {
	a.TranslateXY(v.X, v.Y)
}
*/
