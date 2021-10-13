class Solution {
    public void rotate(int[] nums, int k) {
        reverse(nums,0,nums.length);
        reverse(nums,0,k-1);
        reverse(nums,k,nums.length-1);
    }
    public void reverse(int[] nums,int start,int end){
        int tmp=0;
        while(end>=start){
            tmp = nums[end];
            nums[end--] = nums[start];
            nums[start++] = tmp;
        }
    }
}
