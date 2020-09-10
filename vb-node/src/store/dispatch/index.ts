//store
import { Dispatch } from 'redux';

//inteface
import { IVBAction } from '../action/define';


export const VBDispatch = (action: IVBAction<any>) => (dispatch: Dispatch<IVBAction<any>>, getStore?: () => any) => {
    dispatch(action)
}