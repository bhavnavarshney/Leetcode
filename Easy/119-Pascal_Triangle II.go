func getRow(rowIndex int) []int {
    triangle:=make([][]int,rowIndex+1,rowIndex+1)
    triangle[0]=[]int{1}
    if rowIndex==0{
        return triangle[0]
    }
   

    for i:=1;i<=rowIndex;i++{
        cur_row:=make([]int,i+1)
        prev_row:=triangle[i-1]
        cur_row[0]=1
        cur_row[len(cur_row)-1]=1
        
        for j:=1;j<i;j++{
            cur_row[j] = prev_row[j-1]+prev_row[j]
        }
        triangle[i] = cur_row
    }
    return triangle[rowIndex]
}