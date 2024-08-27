export const loadState = () => {
    const localState = localStorage.getItem("state");
    if (localState === null) {
        return [];
    }
    return JSON.parse(localState);
};

export const saveState = (state: any) => {
    const newState = JSON.stringify(state);
    localStorage.setItem("state", newState);
};