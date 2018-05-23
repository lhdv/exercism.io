// Package space calculates an age based on a given planet
package space

// Planet give name to a planet
type Planet string

// EarthYear means a Earth year in seconds
const EarthYear = 31557600

func getOrbitalYears() map[Planet]float64 {
	return map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
}

// Age should do it's job, well.
func Age(s float64, p Planet) float64 {

	planets := getOrbitalYears()

	return s / (planets[p] * EarthYear)
}
