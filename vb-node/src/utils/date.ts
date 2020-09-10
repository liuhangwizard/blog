

interface IVBUtilFormatDateTable{
    show:number,
    toast:(day:number)=>string
}


const dateTable:IVBUtilFormatDateTable[]=[
    {
        show:365,
        toast:(d)=>Math.floor(d/365)+"年前"
    },
    {
        show:30,
        toast:(d)=>Math.floor(d/30)+"月前"
    },
    {
        show:7,
        toast:(d)=>Math.floor(d/7)+"周前"
    },
    {
        show:0,
        toast:(d)=>d>0?d+"天前":"1天前"
    }
];

export const formatDateTime=(startDate:string):string=>{

    const date=new Date().getTime()- new Date(startDate).getTime();
    const days=Math.floor(date/(24*3600*1000));    



    for(let i of dateTable){
        if(days>=i.show){
            return i.toast(days);
        }
    }
   
}


export const formateDate=(fmt, date)=>{
    let ret;
    const opt = {
        "Y+": date.getFullYear().toString(),        // 年
        "m+": (date.getMonth() + 1).toString(),     // 月
        "d+": date.getDate().toString(),            // 日
        "H+": date.getHours().toString(),           // 时
        "M+": date.getMinutes().toString(),         // 分
        "S+": date.getSeconds().toString()          // 秒

    };
    for (let k in opt) {
        ret = new RegExp("(" + k + ")").exec(fmt);
        if (ret) {
            fmt = fmt.replace(ret[1], (ret[1].length == 1) ? (opt[k]) : (opt[k].padStart(ret[1].length, "0")))
        };
    };
    return fmt;
}