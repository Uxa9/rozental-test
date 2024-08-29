import { Key, useEffect, useState } from "react"
import ElementView from "./elementView"
import { connect, shallowEqual, useDispatch, useSelector } from "react-redux"
import productStoreDispatcher, { DispatcherProps } from "../../common/_productStore/dispatcher"
import { ProductStoreProps, StateProps, productStoreState } from "../../common/_productStore"
import Table from "../../components/table/Table"
import { Product } from "../../common/_productStore/product"
import productStoreReducer from "../../common/_productStore/reducer"
import { ProductReducer } from "../../common/productStore/slice"
import { useAppDispatch } from "../.."
import { RootState } from "../../common/indexStore"
import { fetchPosts, Post } from "../../common/post/slice"
import List from "../../components/list/List"


const TableView: React.FC = () => {
    const {products, post} = useSelector((state: RootState) => state)

    const dispatch = useAppDispatch()
    
    
    useEffect(() => {
        dispatch(fetchPosts())
    }, [])

    useEffect(() => {
        console.log("postList refreshed");
        console.log(post);
        
    }, [post])

    return (
        <>
            <Table elements={products.map((element: Product, index: Key) => {
                return (
                    <ElementView
                        key={index}
                        product={element}
                    />
                )
            })} />
            <List elements={post.postList.map((post: Post, index: Key) => {
                return (
                    <div key={index}>
                        <p>{post.title}</p>
                        <p>{post.body}</p>
                    </div>
                )
            })} />
        </>
    )
}

export default TableView