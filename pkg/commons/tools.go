package commons

import (
	"errors"

	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func StringIsEmpty(e string) bool {
	return len(strings.TrimSpace(e)) == 0
}

func IsEmailValid(e string) bool {
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	email := strings.TrimSpace(e)
	if len(email) < 3 || len(email) > 254 {
		return false
	}
	return pattern.MatchString(email)
}

func EsCuilValido(cuil string) error {
	if len(cuil) != 11 {
		return errors.New(ERROR_CUIL)
	}
	var rv bool
	var verificador int
	resultado := 0
	codes := "5432765432"
	ultimoDigito := cuil[10:11]
	verificador, err := strconv.Atoi(ultimoDigito)
	if err != nil {
		return errors.New(ERROR_CUIL)
	}
	for x := 0; x < 10; x++ {
		digitoValidador, _ := strconv.Atoi(codes[x : x+1])
		digito, _ := strconv.Atoi(cuil[x : x+1])
		digitoValidacion := digitoValidador * digito
		resultado += digitoValidacion
	}
	//resultado = resultado / 11
	resto := resultado % 11
	r2 := 11 - resto
	rv = (r2 == verificador)
	if !rv {
		return errors.New(ERROR_CUIL)
	}
	return nil
}

//////////////////////
// func validarLargoCbu(cbu string) error {
// 	if StringIsEmpity(cbu) {
// 		return fmt.Errorf("cbu está en blanco")
// 	}
// 	if len(cbu) != 22 {
// 		return fmt.Errorf("longitud de cbu no es válido: %d", len(cbu))
// 	}
// 	return nil
// }

// func validarCodigoBanco(codigo string) error {
// 	if len(codigo) != 8 {
// 		return fmt.Errorf("el código de banco es incorrecto")
// 	}
// 	banco := codigo[0:3]
// 	// fmt.Println("banco: " + banco)
// 	digitoVerificador := codigo[3:4]
// 	// fmt.Println("digito verificador: " + digitoVerificador)
// 	sucursal := codigo[4:7]
// 	// fmt.Println("sucursal: " + sucursal)
// 	digitoVerificador2 := codigo[7:8]
// 	// fmt.Println("digito verificador 2: " + digitoVerificador2)

// 	var suma int
// 	var x int

// 	x, _ = strconv.Atoi(banco[0:1])
// 	suma = x * 7
// 	x, _ = strconv.Atoi(banco[1:2])
// 	suma = suma + x
// 	x, _ = strconv.Atoi(banco[2:3])
// 	suma = suma + (x * 3)
// 	x, _ = strconv.Atoi(digitoVerificador)
// 	suma = suma + (x * 9)
// 	x, _ = strconv.Atoi(sucursal[0:1])
// 	suma = suma + (x * 7)
// 	x, _ = strconv.Atoi(sucursal[1:2])
// 	suma = suma + x
// 	x, _ = strconv.Atoi(sucursal[2:3])
// 	suma = suma + (x * 3)

// 	diferencia := 10 - (suma % 10)
// 	digito, _ := strconv.Atoi(digitoVerificador2)
// 	if diferencia != digito {
// 		return fmt.Errorf("codigo de banco inválido")
// 	}
// 	return nil
// }

// func validarCuenta(cuenta string) error {
// 	if len(cuenta) != 14 {
// 		return fmt.Errorf("logitud de cuenta inválido: %d", len(cuenta))
// 	}
// 	digitoVerificador, _ := strconv.Atoi(cuenta[13:14])

// 	var suma int
// 	var x int

// 	x, _ = strconv.Atoi(cuenta[0:1])
// 	suma = x * 3
// 	x, _ = strconv.Atoi(cuenta[1:2])
// 	suma = suma + (x * 9)
// 	x, _ = strconv.Atoi(cuenta[2:3])
// 	suma = suma + (x * 7)
// 	x, _ = strconv.Atoi(cuenta[3:4])
// 	suma = suma + x
// 	x, _ = strconv.Atoi(cuenta[4:5])
// 	suma = suma + (x * 3)
// 	x, _ = strconv.Atoi(cuenta[5:6])
// 	suma = suma + (x * 9)
// 	x, _ = strconv.Atoi(cuenta[6:7])
// 	suma = suma + (x * 7)
// 	x, _ = strconv.Atoi(cuenta[7:8])
// 	suma = suma + (x * 1)
// 	x, _ = strconv.Atoi(cuenta[8:9])
// 	suma = suma + (x * 3)
// 	x, _ = strconv.Atoi(cuenta[9:10])
// 	suma = suma + (x * 9)
// 	x, _ = strconv.Atoi(cuenta[10:11])
// 	suma = suma + (x * 7)
// 	x, _ = strconv.Atoi(cuenta[11:12])
// 	suma = suma + (x * 1)
// 	x, _ = strconv.Atoi(cuenta[12:13])
// 	suma = suma + (x * 3)

// 	diferencia := 10 - (suma % 10)

// 	if diferencia != digitoVerificador {
// 		return fmt.Errorf("error en cuenta bancaria")
// 	}

// 	return nil
// }

// func ValidarCBU(cbu string) error {
// 	err := validarLargoCbu(cbu)
// 	if err != nil {
// 		return err
// 	}
// 	err = validarCodigoBanco(cbu[0:8])
// 	if err != nil {
// 		return err
// 	}
// 	err = validarCuenta(cbu[8:22])
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

//// validara tarjeta de credito

// func obtenerLongitud(nro string) (int, error) {
// 	if StringIsEmpity(nro) && len(nro) >= 12 && len(nro) <= 16 {
// 		return 0, fmt.Errorf("longitud no valida %v", nro)
// 	} else {
// 		return int(len(nro)), nil
// 	}
// }
///////////////////////////////////////////////////////////////////////

// func obtenerLongitud(nro string) (int, error) {
// 	if StringIsEmpity(nro) && len(nro) >= 12 && len(nro) <= 16 {
// 		return 0, fmt.Errorf("longitud no valida %v", nro)
// 	} else {
// 		return int(len(nro)), nil
// 	}
// }

// func SumarDigitos(digito string) (int, error) {
// 	if len(digito) == 1 {
// 		suma, _ := strconv.Atoi(digito)
// 		return suma, nil
// 	} else if len(digito) == 2 {
// 		valor1, _ := strconv.Atoi(digito[0:1])
// 		valor2, _ := strconv.Atoi(digito[1:2])
// 		suma := valor1 + valor2
// 		return suma, nil
// 	} else {
// 		return 0, fmt.Errorf("longitud no valida %v", digito)
// 	}
// }

// func DuplicarValor(digito string) string {
// 	x, _ := strconv.Atoi(digito)
// 	z := x * 2
// 	y := strconv.Itoa(z)
// 	return y
// }

// func validarTarjeta(nroCard string) int {
// 	var suma int
// 	var sumaDoble int
// 	par := len(nroCard) % 2
// 	if par == 0 {
// 		for i := len(nroCard); i > 0; i-- {
// 			if i != len(nroCard) {
// 				if (i % 2) == 1 {
// 					digitoDoble := DuplicarValor(nroCard[i-1 : i])
// 					sumadigito, _ := SumarDigitos(digitoDoble)
// 					sumaDoble = sumaDoble + sumadigito
// 					//   fmt.Printf("%v - %v - %v - %v -e %v\n", i ,i%2, nroCard[i-1:i], digitoDoble, sumadigito  )
// 				} else {
// 					x, _ := strconv.Atoi(nroCard[i-1 : i])
// 					suma = suma + x
// 				}
// 			} else {
// 				x, _ := strconv.Atoi(nroCard[i-1 : i])
// 				suma = suma + x
// 			}
// 		}
// 		return suma + sumaDoble
// 		//   fmt.Printf("%v - %v\n", suma, sumaDoble)
// 		//   fmt.Printf("%v \n", suma + sumaDoble)
// 	} else {
// 		for i := len(nroCard); i >= 1; i-- {
// 			if i != len(nroCard) {
// 				if (i % 2) == 0 {
// 					digitoDoble := DuplicarValor(nroCard[i-1 : i])
// 					sumadigito, _ := SumarDigitos(digitoDoble)
// 					sumaDoble = sumaDoble + sumadigito
// 					//fmt.Printf("%v - %v - %v\n", i, i%2, nroCard[i-1:i])
// 				} else {
// 					x, _ := strconv.Atoi(nroCard[i-1 : i])
// 					suma = suma + x
// 				}
// 			} else {
// 				x, _ := strconv.Atoi(nroCard[i-1 : i])
// 				suma = suma + x
// 			}
// 		}
// 		// fmt.Printf("%v - %v\n", suma, sumaDoble)
// 		// fmt.Printf("%v \n", suma+sumaDoble)
// 		return suma + sumaDoble
// 	}
// }

// ChequearTarjeta es la funcio que se debe llamar para validar el formato de una tarjeta
// puede retornanar uno de los siguientes valores
// True: si el formato de tarjeta es valido
// False: si el formato de tarjeta no es valido
// func ChequearTarjeta(valorCheck string) bool {
// 	longitudValida, _ := obtenerLongitud(valorCheck)
// 	if longitudValida != 0 {
// 		valorVerificador := validarTarjeta(valorCheck)
// 		valorModulo := valorVerificador % 10
// 		if valorModulo == 0 {
// 			return true
// 		} else {
// 			return false
// 		}
// 	}
// 	return false
// }

////////////////////////////////////////////////////////////////////////

func Difference(slice1 []string, slice2 []string) []string {
	var diff []string

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				if len(s1) > 0 {
					diff = append(diff, s1)
				}
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}

/*
	funciones que se pueden utilizar para quitar caracteres especiales en una cadena
*/
func SpaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func StripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			// if the character is a space, drop it
			return -1
		}
		// else keep it in the string
		return r
	}, str)
}

/*
	remplazar un caracter por otro
*/
func ReplaceCharacters(str, valorBuscar, valorReemplazar string) string {
	resultadoString := strings.Replace(str, valorBuscar, valorReemplazar, -1)
	return resultadoString
}
