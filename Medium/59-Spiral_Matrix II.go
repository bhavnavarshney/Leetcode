func generateMatrix(n int) [][]int {
    var a [][]int
    for i:=0;i<n;i++{
        temp := make([]int,n)
        a = append(a,temp)
    }
    //fmt.Println(a)
    r1:=0; r2:=n-1
    c1:=0; c2:=n-1
    k:=1
    
    for r1<=r2 && c1<=c2{
        for i:=c1;i<=c2;i+=1{
            a[r1][i]=k
            k++
        }
        for i:=r1+1;i<=r2;i++{
            a[i][c2] = k
            k++
        }
        if r1<r2 && c1<c2{
            for i:=c2-1;i>=c1;i--{
                a[r2][i]=k
                k++
            }
            for i:=r2-1;i>r1;i--{
                a[i][c1] = k
                k++
            }
        }
        r1++; c1++; r2--; c2--;
    }
    
    return a
}