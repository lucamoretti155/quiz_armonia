# quiz_armonia

Il programma ha come scopo quello di proporre domande all'utente sull'armonia musicale

Il programma all'inizio esegue un caricamento tramite 3 file _*.txt_ presenti nella stessa cartella del file _main.go_ che contengono tutte le info per popolare delle variabili all'interno del programma.
Le domande possono essere di 5 tipoligie differenti su tematiche specifiche proposte in maniera aleatoria, ma non esiste un numero finito di domande. L'utente può continuare il quiz per tutto il tempo che vuole. Ad ogni domanda verrà chiesto se questi vuole continuare o meno. Qualora l'utente decida di non continuare il programma termina salutando e stampando il numero di risposte esatte fornite e il relativo numero di domande totali.


Il programma non effettua molti controlli circa i dati di input se non la coerenza con la risposta fornita. Quindi l'assumption è che l'utente sia collaborativo e rispetti i suggerimenti di modalità di risposta fornita ad ogni domanda. Il check delle risposte non è case sensitive.


Esempi di esecuzione:

Domanda tipologia 1

    *** Benvenuto! Per iniziare premi un tasto qualsiasi ***
    s
    *** Domanda numero 1 ***
    Quante alterazioni ha la scala di do# minore? (e.g. ==> 3 #, se zero allora 0 #)
    4 #
    Risposta esatta!
    
    *** Continuare?(y/n) ***
    y
    *** Domanda numero 2 ***
    Quante alterazioni ha la scala di Fa maggiore? (e.g. ==> 3 #, se zero allora 0 #)
    2 b
    Risposta errata. La risposta corretta era:  1 b

Domanda tipologia 2

    *** Benvenuto! Per iniziare premi un tasto qualsiasi ***
    x
    *** Domanda numero 1 ***
    Quali sono le note alterate nella scala di re# minore? (e.g. ==> Do# Re# ...; se la nessuna allora scrivere ==> nessuna)
    fa# do# sol# re# la# mi#
    Risposta esatta!

    *** Continuare?(y/n) ***
    y
    *** Domanda numero 2 ***
    Quali sono le note alterate nella scala di Mi maggiore? (e.g. ==> Do# Re# ...; se la nessuna allora scrivere ==> nessuna)
    do#
    Risposta errata. La risposta corretta era:  [Fa# Do# Sol# Re#]

Domanda tipologia 3

    *** Benvenuto! Per iniziare premi un tasto qualsiasi ***
    x
    *** Domanda numero 1 ***
    Qual'è la relativa scala minore di Dob maggiore? (e.g. Sib minore)
    lab miNORE
    Risposta esatta!

    *** Continuare?(y/n) ***
    y
    *** Domanda numero 2 ***
    Qual'è la relativa scala minore di La maggiore? (e.g. Sib minore)
    solb minore
    Risposta errata. La risposta corretta era:  fa# minore
    
Domanda tipologia 4

    *** Benvenuto! Per iniziare premi un tasto qualsiasi ***
    x
    *** Domanda numero 1 ***
    In che tonalità maggiore si trova l'accordo Sol-7
    fa maggiore
    Risposta esatta!
    *** Info extra ***
    L'accordo Sol-7 è ii grado in Fa maggiore
    L'accordo Sol-7 è iii grado in Mib maggiore
    L'accordo Sol-7 è vi grado in Sib maggiore

    *** Continuare?(y/n) ***
    y
    *** Domanda numero 2 ***
    In che tonalità maggiore si trova l'accordo Re#m7b5
    do maggiore
    Risposta errata.
    *** Info extra ***
    L'accordo Re#m7b5 è vii grado in Mi maggiore
    
 Domanda tipologia 5

    *** Benvenuto! Per iniziare premi un tasto qualsiasi ***
    x
    *** Domanda numero 1 ***
    Qual'è l'intervallo fra le note DO - MI? (e.g. seconda maggiore)
    terza maggiore
    Risposta esatta!

    *** Continuare?(y/n) ***
    y
    *** Domanda numero 2 ***
    Qual'è l'intervallo fra le note RE - SI? (e.g. seconda maggiore)
    sesta minore
    Risposta errata. La risposta corretta era: sesta maggiore
    
    
  
