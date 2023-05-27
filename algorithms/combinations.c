unsigned char NextSequenceFastest(unsigned char * array, unsigned char base) {
	int i;

	for (i = 0; (*(array+i)) < base; i++) {
		(*(array+i))++;

		if (*(array+i) < base) {
			return 1;
		}
		(*(array+i)) = 0;
	}

	return 0;
}

unsigned long NextSequenceFastest_Count_1048575_batch(int repeatNTimes) {

    int n;
   	unsigned long i = 0;


    for (n=0;n<repeatNTimes;n++){
        unsigned char array[] = {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5};

    	while(NextSequenceFastest(array, 4)) {
		    i++;
	    }

    }

    return i;

}