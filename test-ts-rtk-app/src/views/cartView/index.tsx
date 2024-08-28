import styles from "./CartView.module.scss";
import { useNavigate } from "react-router-dom";
import { GetRandomSusById } from "../../randomSusAssert";
import { CartStoreAndOwnProps, CartStoreProps, cartStoreState, cartStoreStateAndProps, StateProps } from "../../common/cart";
import CartDispatcher, { DispatcherProps as CartDispatcherProps } from "../../common/cart/dispatcher";
import { DispatcherProps as ProductDispatcherProps } from "../../common/productStore/dispatcher";
import { connect, useDispatch, useSelector } from "react-redux";
import { StoreProps, StoreState, StoreDispathcer, CommonStore, StoreDispatcherProps } from "../../common/indexStore";
import { useState, useEffect } from "react";
import { Product } from "../../common/productStore/product";
import List from "../../components/list/List";
import ElementView from "./elementView";
import Button from "../../components/button/Button";


const CartView = () => {

    const dispatch = useDispatch()

    const { cart, products } = useSelector((state: any) => state)
    const [elementList, setElementList] = useState<any[]>([])

    const [positionsMap, setPositionsMap] = useState<Map<number, number>>(new Map)

    useEffect(() => {
        const map = cart.reduce((acc: any, e: any) => acc.set(e, (acc.get(e) || 0) + 1), new Map())
        setPositionsMap(map)        
    }, [ ,cart])

    useEffect(() => {
        let elList: any[] = [];

        Array.from(positionsMap.keys()).forEach((key) => {
            const product = products.find((el: { id: number; }) => el.id === key);
            const amount = positionsMap.get(key);

            if (product && amount) {
                elList.push({
                    product,
                    amount
                });
            }
        })

        setElementList(elList);
    }, [positionsMap])

    const calculateTotalSum = () => {
        return elementList.reduce((acc, current) => {
            return acc + (current.product.price * current.amount);
        }, 0);
    }

    const handleWipeButton = () => {
        dispatch(CartDispatcher.wipeCart())
    }


    return (
        <>
            {cart.length ?
                <>
                    <List elements={elementList.map((item, index) => {
                        return (
                            <ElementView
                                key={index}
                                product={item.product}
                                amount={item.amount}
                            />
                        )
                    })
                    } />
                    <p>
                        Итого: {calculateTotalSum()} ₽
                    </p>
                    <div className={styles["wipe-button"]}>
                        <Button onClick={handleWipeButton} text="Очистить корзину" />
                    </div>
                </>
                :
                <span>Корзина пуста</span>
            }
        </>
    )
}

export default CartView