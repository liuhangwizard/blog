//history
import { createBrowserHistory} from 'history';

const VBHistory=typeof document==="object"?createBrowserHistory():undefined;


export default VBHistory;