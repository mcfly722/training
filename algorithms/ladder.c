signed char NextLadder(unsigned char * array,unsigned char length, unsigned char maxHeight) {

    // get last correct modification position
	signed char pos = length - 1;
	signed char max = maxHeight - pos;

	for (; pos > -1 && *(array+pos) >= max+pos; pos--) {
	}

	if (pos == -1) {
		return -1;
	}

	signed char r = pos;

	// increment it
	(*(array+pos))++;
	pos++;

	// reset tail
	for (; pos < length; pos++) {
		*(array+pos) = *(array+pos-1) + 1;
	}

	return r;
}

unsigned long NextLadder_Count_352715_batch(int repeatNTimes) {
    int n;
   	unsigned long i = 0;
    for (n=0;n<repeatNTimes;n++){
        unsigned char ladder[] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
    	while(NextLadder(ladder, 10, 21)!=-1) {
		    i++;
	    }
    }
    return i;
}