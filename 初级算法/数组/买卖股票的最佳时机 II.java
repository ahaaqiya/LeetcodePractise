class Solution {
    public int maxProfit(int[] prices) {
		int lowest=0,highest=0;
		int index=0,sum=0;
		if(prices.length<=1)
			return sum;
		while(index<prices.length-1) {
			while(index<prices.length-1 && prices[index]>=prices[index+1])
				index++;
			lowest = prices[index];
			while(index<prices.length-1 && prices[index]<=prices[index+1])
				index++;
			highest = prices[index];
			sum += (highest-lowest);	
		}
		return sum;
    }
}
