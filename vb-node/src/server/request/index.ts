//axios
import Axios, { AxiosInstance } from 'axios';

//config
import { serverConfig } from '../config';

//axios
//client
export const clientAxios=Axios.create({
    baseURL:serverConfig.clientBaseURL,
    timeout:serverConfig.axiosTimeout,
})
//server
export const serverAxios=Axios.create({
    baseURL:serverConfig.serverBaseURL,
    timeout:serverConfig.axiosTimeout
})

//target
//const isServer=typeof document==='undefined';
//export const getRequest=():AxiosInstance=> isServer?serverAxios:clientAxios;

