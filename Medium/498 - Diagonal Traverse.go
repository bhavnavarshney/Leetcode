func findDiagonalOrder(matrix [][]int) []int {
    if len(matrix)<1 || len(matrix[0])<1{
        return make([]int,0)
    }
    
    res := make([]int,len(matrix[0])*len(matrix))
    i:=0; j:=0; index:=0
    goUp(matrix,res,i,j,index)
    return res
}

func goUp(matrix [][]int, res []int, i int, j int,index int){
    res[index] = matrix[i][j]
    for i!=0 && j!=len(matrix[0])-1{
        i--; j++; index++;
        res[index] = matrix[i][j]
    }
    
    if j==len(matrix[0])-1{
        i+=1
    }else{
        j+=1
    }
  
    if index<(len(matrix)*len(matrix[0])-1){
        goDown(matrix,res,i,j,index+1)
    }
}

func goDown(matrix [][]int,res []int,i int,j int,index int){
    res[index] = matrix[i][j]
    
    for j!=0 && i!=len(matrix)-1{
        j--; i++; index+=1;
        res[index] = matrix[i][j]
    }
    
    if i==len(matrix)-1 {
        j+=1
    }else{
        i+=1
    }
    
    if index<(len(matrix)*len(matrix[0])-1){
        goUp(matrix,res,i,j,index+1)
    }
}