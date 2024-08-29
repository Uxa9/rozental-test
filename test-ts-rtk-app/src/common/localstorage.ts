import { RootState } from "./indexStore";


export const loadState = (storeName: string) => {
    const localState = localStorage.getItem("state");
    if (localState === null) {
        return [];
    }

    const parsedState = JSON.parse(localState);
    
    return parsedState[storeName]
};

export const saveState = (state: any) => {
    const newState = JSON.stringify(state);
    localStorage.setItem("state", newState);
};