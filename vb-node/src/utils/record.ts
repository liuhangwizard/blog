
export const formatRecord=(record:number):string=>{
    if(record>=100){
        return Number(record/1000).toFixed(1)+"k"
    }
    return "0.1k"
}