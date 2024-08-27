import styles from "./CartView.module.scss";
import { useNavigate } from "react-router-dom";
import { GetRandomSusById } from "../../randomSusAssert";
import { CartStoreAndOwnProps, CartStoreProps, cartStoreState, cartStoreStateAndProps, StateProps } from "../../common/cart";
import CartDispatcher, { DispatcherProps as CartDispatcherProps } from "../../common/cart/dispatcher";
import { DispatcherProps as ProductDispatcherProps } from "../../common/productStore/dispatcher";
import { connect } from "react-redux";
import { StoreProps, StoreState, StoreDispathcer, CommonStore, StoreDispatcherProps } from "../../common/indexStore";
import { useState, useEffect } from "react";
import { Product } from "../../common/productStore/product";
import List from "../../components/list/List";
import ElementView from "./elementView";
import Button from "../../components/button/Button";


const CartView: React.FC<CommonStore> = (props) => {

    const [elements, setElements] = useState<Product[]>([])
    const [positionsMap, setPositionsMap] = useState<Map<number, number>>(new Map)
    const [elementList, setElementList] = useState<any[]>([])

    useEffect(() => {
        setElements(props.elements)

        const map = props.cart.reduce((acc, e) => acc.set(e, (acc.get(e) || 0) + 1), new Map())
        setPositionsMap(map)
        
    }, [ , props.cart])

    useEffect(() => {
        let elList: any[] = [];

        Array.from(positionsMap.keys()).forEach((key) => {
            const product = elements.find(el => el.id === key);
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


    return (
        <>
            {props.cart.length ?
                <>
                    <List elements={elementList.map((item, index) => {
                        return (
                            <ElementView
                                key={index}
                                own={item}
                            />
                        )
                    })
                    } />
                    <p>
                        Итого: {calculateTotalSum()} ₽
                    </p>
                    <div className={styles["wipe-button"]}>
                        <Button onClick={props.wipeCart} text="Очистить корзину" />
                    </div>
                </>
                :
                <span>Корзина пуста</span>
            }
        </>
    )
}

export default connect<StoreProps, StoreDispatcherProps>(
    StoreState,
    StoreDispathcer
)(CartView)