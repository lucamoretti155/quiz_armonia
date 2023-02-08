package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type InfoAccordo struct {
	tonalitàMaggiore string
	tonalitàMinore   string
	gradoInMaggiore  string
	gradoInMinore    string
}

type Accordo []InfoAccordo

type Scala struct {
	maggiore          string
	relativaMinore    string
	numeroAlterazioni int
	tipoAlterazioni   string
	noteAlterate      []string
	noteScala         [7]string
}

func main() {
	f, err := os.Open("intervalli.txt")
	if err != nil {
		fmt.Println("errore nell'apertura del file")
		return
	}
	defer f.Close()
	noteIntervalli, mappaIntervalli, err := LeggiIntervalli(f)
	fileScale, err := os.Open("scale.txt")
	if err != nil {
		fmt.Println("errore nell'apertura del file")
		return
	}
	defer fileScale.Close()
	fileAccordi, err := os.Open("accordi.txt")
	if err != nil {
		fmt.Println("errore nell'apertura del file")
		return
	}
	defer fileAccordi.Close()
	var scale []Scala
	scale, err = LeggiScale(fileScale)
	if err != nil {
		return
	}
	accordi := make(map[string]Accordo)
	accordi, err = LeggiAccordi(fileAccordi)
	if err != nil {
		return
	}

	//check manuale della lettura
	/*
		fOutScale, err := os.Create("check_scale.txt")
		if err != nil {
			fmt.Println("errore nella creazione del file")
			return
		}
		defer fOutScale.Close()
		for k, v := range scale {
			_, err = fmt.Fprintf(fOutScale, "%v: %v\n", k, v)
			if err != nil {
				fmt.Println("errore nella scrittura del file")
				return
			}
		}
		fOutAccordi, err := os.Create("check_accordi.txt")
		if err != nil {
			fmt.Println("errore nella creazione del file")
			return
		}
		defer fOutAccordi.Close()
		for k, v := range accordi {
			_, err = fmt.Fprintf(fOutAccordi, "%v: %v\n", k, v)
			if err != nil {
				fmt.Println("errore nella scrittura del file")
				return
			}
		}

	*/

	var tipoDomanda int
	var s string
	var contaDomande, contaEsatte int = 0, 0
	fmt.Println("*** Benvenuto! Per iniziare premi un tasto qualsiasi ***")
	fmt.Scan(&s)
	for {
		rand.Seed(int64(time.Now().Nanosecond()))
		tipoDomanda = rand.Intn(5)

		contaDomande++
		fmt.Println("*** Domanda numero", contaDomande, "***")
		switch tipoDomanda {
		case 0:
			contaEsatte += QuanteAlterazioni(scale)
		case 1:
			contaEsatte += QualiAlterazioni(scale)
		case 2:
			contaEsatte += RelativaMinore(scale)
		case 3:
			contaEsatte += DoveTonicizza(accordi)
		case 4:
			contaEsatte += Intervalli(mappaIntervalli, noteIntervalli)
		default:
			continue
		}
		fmt.Println()
		fmt.Println("*** Continuare?(y/n) ***")
		fmt.Scan(&s)
		if s != "y" {
			break
		}
	}
	fmt.Println()
	fmt.Println("Risposte esatte =", contaEsatte, " Domande in totale =", contaDomande)
	fmt.Println("*** CIAO! ***")
}

func LeggiScale(fileScale *os.File) (scale []Scala, err error) {
	scanner := bufio.NewScanner(fileScale)
	var scala Scala
	scanner.Scan()
	for scanner.Scan() {
		riga := strings.Split(scanner.Text(), ";")
		scala.maggiore = riga[0]
		scala.relativaMinore = riga[1]
		scala.numeroAlterazioni, _ = strconv.Atoi(riga[2])
		scala.tipoAlterazioni = riga[3]
		scala.noteAlterate = []string{}
		for i := 0; i < scala.numeroAlterazioni; i++ {
			scala.noteAlterate = append(scala.noteAlterate, riga[i+4])
		}
		for i := 0; i < 7; i++ {
			scala.noteScala[i] = riga[i+4+7]
		}
		scale = append(scale, scala)
	}
	if err = scanner.Err(); err != nil {
		fmt.Printf("Error while reading the file! %v\n", err)
		return
	}
	return
}

func LeggiAccordi(fileAccordi *os.File) (accordi map[string]Accordo, err error) {
	scanner := bufio.NewScanner(fileAccordi)
	var accordo InfoAccordo
	accordi = make(map[string]Accordo)
	scanner.Scan()
	for scanner.Scan() {
		riga := strings.Split(scanner.Text(), ";")
		accordo.tonalitàMaggiore = riga[1]
		accordo.tonalitàMinore = riga[2]
		accordo.gradoInMaggiore = riga[3]
		accordo.gradoInMinore = riga[4]
		accordi[riga[0]] = append(accordi[riga[0]], accordo)
	}
	if err = scanner.Err(); err != nil {
		fmt.Printf("Error while reading the file! %v\n", err)
		return
	}
	return
}

func QuanteAlterazioni(scale []Scala) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	n1 := rand.Intn(15)
	n2 := rand.Intn(2)
	scalaEstratta := scale[n1]
	var nomeScala string
	if n2 == 0 {
		nomeScala = scalaEstratta.maggiore
	} else {
		nomeScala = scalaEstratta.relativaMinore
	}
	alterazioni := scalaEstratta.numeroAlterazioni
	tipoAlterazioni := scalaEstratta.tipoAlterazioni
	fmt.Printf("Quante alterazioni ha la scala di %v? (e.g. ==> 3 #, se zero allora 0 #)\n", nomeScala)
	var rispostaNumero int
	var rispostaTipo string
	fmt.Scan(&rispostaNumero, &rispostaTipo)
	if rispostaNumero == alterazioni && rispostaTipo == tipoAlterazioni {
		fmt.Println("Risposta esatta!")
		return 1
	} else {
		fmt.Println("Risposta errata. La risposta corretta era: ", alterazioni, tipoAlterazioni)
		return 0
	}
}

func QualiAlterazioni(scale []Scala) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	n1 := rand.Intn(15)
	n2 := rand.Intn(2)
	scalaEstratta := scale[n1]
	var nomeScala string
	if n2 == 0 {
		nomeScala = scalaEstratta.maggiore
	} else {
		nomeScala = scalaEstratta.relativaMinore
	}
	fmt.Printf("Quali sono le note alterate nella scala di %v? (e.g. ==> Do# Re# ...; se la nessuna allora scrivere ==> nessuna)\n", nomeScala)
	alterazioni := scalaEstratta.numeroAlterazioni
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	risposta := strings.Split(strings.ToLower(scanner.Text()), " ")
	contaAlterazioni := 0
	for _, r := range risposta {
		for _, note := range scalaEstratta.noteAlterate {
			if strings.ToLower(note) == r {
				contaAlterazioni++
			}
		}
	}
	if contaAlterazioni == alterazioni {
		fmt.Println("Risposta esatta!")
		return 1
	} else {
		fmt.Println("Risposta errata. La risposta corretta era: ", scalaEstratta.noteAlterate)
		return 0
	}
}

func RelativaMinore(scale []Scala) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	n1 := rand.Intn(15)
	scalaEstratta := scale[n1]
	maggiore := scalaEstratta.maggiore
	minore := scalaEstratta.relativaMinore
	fmt.Printf("Qual'è la relativa scala minore di %v? (e.g. Sib minore)\n", maggiore)
	var risposta string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	risposta = strings.ToLower(scanner.Text())
	if strings.ToLower(risposta) == strings.ToLower(minore) {
		fmt.Println("Risposta esatta!")
		return 1
	} else {
		fmt.Println("Risposta errata. La risposta corretta era: ", minore)
		return 0
	}
}

func DoveTonicizza(accordi map[string]Accordo) int {
	chiave := AccordoCasuale(accordi)
	var AccordoEstratto Accordo
	AccordoEstratto = accordi[chiave]
	var risposta string
	fmt.Printf("In che tonalità maggiore si trova l'accordo %v?\n", chiave)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	risposta = strings.ToLower(scanner.Text())
	ok := false
	var infoAggiuntiva string = "*** Info extra ***\n"
	for _, info := range AccordoEstratto {
		x := strings.Split(info.tonalitàMaggiore, " ")[0]
		if strings.ToLower(x) == risposta {
			ok = true
		}
		infoAggiuntiva += fmt.Sprintf("L'accordo %v è %v grado in %v\n", chiave, info.gradoInMaggiore, info.tonalitàMaggiore)
	}
	var conta int
	if ok {
		fmt.Println("Risposta esatta!")
		conta = 1
	} else {
		fmt.Println("Risposta errata.")
		conta = 0
	}
	fmt.Print(infoAggiuntiva)
	return conta
}

func AccordoCasuale(accordi map[string]Accordo) string {
	rand.Seed(int64(time.Now().Nanosecond()))
	n := rand.Intn(64)
	var sl []string
	for k, _ := range accordi {
		sl = append(sl, k)
	}
	return sl[n]
}

func LeggiIntervalli(file *os.File) (note []string, m map[string]string, err error) {
	scanner := bufio.NewScanner(file)
	m = make(map[string]string)
	for scanner.Scan() {
		riga := strings.Split(scanner.Text(), ";")
		note = append(note, riga[0]+" - "+riga[1])
		m[riga[0]+" - "+riga[1]] = riga[2]
	}
	if err = scanner.Err(); err != nil {
		fmt.Printf("Error while reading the file! %v\n", err)
		return
	}
	return
}

func Intervalli(m map[string]string, note []string) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	i := rand.Intn(len(note) + 1)
	fmt.Printf("Qual'è l'intervallo fra le note %v? (e.g. seconda maggiore)\n", note[i])
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	risposta := strings.ToLower(scanner.Text())
	if strings.ToLower(risposta) == m[note[i]] {
		fmt.Println("\nRisposta esatta!")
		return 1
	} else {
		fmt.Println("\nRisposta errata. La risposta corretta era: ", m[note[i]])
		return 0
	}
}
