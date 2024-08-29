import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

export type Post = {
    userId: number,
    id: number,
    title: string,
    body: string
}

type InitialState = {
    postList: Post[],
    isLoading: boolean,
    counter: number,
    error: string | null
}

const initialState: InitialState = {
    postList: [],
    isLoading: false,
    counter: 0,
    error: null
}

export const fetchPosts = createAsyncThunk('posts/fetchData', async () => {
    
    try {
        

        const data = await fetch('https://jsonplaceholder.typicode.com/posts')
        
        if (data.status !== 200) throw new Error("Error")
        
        return await data.json()

    } catch (error) {
        console.error(error)
        return []
    }

})


const postSlice = createSlice({
    name: 'posts',
    initialState,
    reducers: {
    },
    extraReducers(builder) {
        builder
        .addCase(fetchPosts.pending, (state) => {
            state.postList = []
            state.isLoading = true
            state.error = null
        })
        .addCase(fetchPosts.fulfilled, (state, action) => {
            state.postList = action.payload
            state.isLoading = false
        })
        .addCase(fetchPosts.rejected, (state) => {
            state.postList = []
            state.error = "Error during fetching data from remote"
        })
    },
})

export const PostReducer = postSlice.reducer