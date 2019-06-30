func isAnagram(s string, t string) bool {
    var s1,s2 [26]int
    
    if(len(s)!=len(t)){
        return false
    }
    for i:=0;i<len(s);i++ {
        s1[s[i]-'a']++
        s2[t[i]-'a']++
    }
    
    if s1!=s2{
        return false
    }
    return true
} 