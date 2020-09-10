export const convertContent=(content:string)=>content.replace(/</g,'').replace(/>/g,'')

export const convertServer=(params:any)=>{
    let newParams={};
    Object.keys(params).forEach((key,i)=>{
        newParams={
            ...newParams,
            [key]:convertContent(params[key])
        }
    })

    return newParams;
}