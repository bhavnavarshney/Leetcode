func generate(numRows int) [][]int {
    
    triangle := make([][]int,numRows,numRows)
    
    if numRows==0 {
        return triangle
    }
    
    triangle[0] = []int{1}
    for i:=1;i<numRows;i++{
        cur_row := make([]int,i+1)
        prev_row := triangle[i-1]
        cur_row[0]=1
        cur_row[len(cur_row)-1]=1
        //fmt.Println("i= ",i,cur_row)
        for j:=1;j<i;j++{
            cur_row[j] = prev_row[j-1]+prev_row[j]
        }
        triangle[i] = cur_row
    }
    
    return triangle
}