class Solution {
    public int removeDuplicates(int[] nums) {
		int left=0;
		for(int right=1;right<nums.length;right++) {
			//left随right向右移动
			if(nums[left]!=nums[right]){
                nums[++left]=nums[right];
            }
					
		}
    //left从0开始，所以返回长度时需要+1
		return ++left;
    }
}
