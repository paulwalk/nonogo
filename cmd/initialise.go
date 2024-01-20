package cmd

func initialiseApplication() {
	var err error
	log, err = ConfigureZapSugarLogger(debug)
	if err != nil {
		log.Fatal("Unable to initialise logging, halting")
	}
}
