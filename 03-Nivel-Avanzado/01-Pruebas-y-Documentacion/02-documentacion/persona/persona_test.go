package persona

import (
	"fmt"
	"strings"
	"testing"
)

func TestNueva(t *testing.T) {
    casos := []struct {
        nombre          string
        edad           int
        deberíaTenerError bool
    }{
        {"Juan", 25, false},
        {"", 25, true},
        {"María", -1, true},
        {"Pedro", 200, true},
    }

    for _, caso := range casos {
        t.Run(fmt.Sprintf("nombre=%s,edad=%d", caso.nombre, caso.edad), func(t *testing.T) {
            p, err := Nueva(caso.nombre, caso.edad)
            if caso.deberíaTenerError {
                if err == nil {
                    t.Error("Se esperaba un error pero no se obtuvo ninguno")
                }
            } else {
                if err != nil {
                    t.Errorf("No se esperaba error pero se obtuvo: %v", err)
                }
                if p.Nombre != caso.nombre {
                    t.Errorf("Nombre esperado %s, se obtuvo %s", caso.nombre, p.Nombre)
                }
                if p.Edad != caso.edad {
                    t.Errorf("Edad esperada %d, se obtuvo %d", caso.edad, p.Edad)
                }
            }
        })
    }
}

func TestSaludar(t *testing.T) {
    p := &Persona{Nombre: "Ana", Edad: 30}
    saludo := p.Saludar()
    if !strings.Contains(saludo, p.Nombre) {
        t.Errorf("El saludo debería contener el nombre '%s'", p.Nombre)
    }
    if !strings.Contains(saludo, fmt.Sprint(p.Edad)) {
        t.Errorf("El saludo debería contener la edad '%d'", p.Edad)
    }
}