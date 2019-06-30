func spiralOrder(matrix [][]int) []int {
    if len(matrix)==0 {
        return []int{}
    }
    
    r1:=0
    c1:=0
    r2:=len(matrix)-1
    c2:=len(matrix[0])-1
    ans := make([]int,len(matrix)*len(matrix[0]))
    k:=0
    
    for i:=c1;i<=c2;i++{
        ans[k] = matrix[r1][i]
        k++
    }
    for i:=r1+1;i<=r2;i++{
        ans[k] = matrix[i][c2]
        k++
    }
    if r1<r2 && c1<c2 {
        for i:=c2-1;i>=r2;i--{
            ans[k] = matrix[r2][i]
            k+=1
        }
        for i:=r2-1;i>=r1+1;i--{
            ans[k] = matrix[i][c1]
            k+=1
        }
    }
    r1+=1
    c1+=1
    r2-=1
    c2-=1
    
    return ans
}