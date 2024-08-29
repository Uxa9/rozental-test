import styles from "./CartView.module.scss";
import {logo} from "../../icons/index"
import { Black } from "../../components/images";
import { useNavigate } from "react-router-dom";
import Button from "../../components/button/Button";
import { Product } from "../../common/_productStore/product";
import { GetRandomSusById } from "../../randomSusAssert";
import { CartStoreAndOwnProps, CartStoreProps, cartStoreState, cartStoreStateAndProps, CartViewProps, StateProps } from "../../common/_cart";
import CartDispatcher, { DispatcherProps } from "../../common/_cart/dispatcher";
import { connect, useDispatch } from "react-redux";
import { useAppDispatch } from "../..";
import { removeFromCart } from "../../common/cart/slice";

type Props = {
    product: Product,
    amount: number
}

const ElementView: React.FC<Props> = (props) => {    

    const dispatch = useAppDispatch()

    const handleButtonClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.stopPropagation()
        dispatch(removeFromCart(props.product.id))
    }

    return (
        <div className={styles["element-view"]}>
            <div>
                <img src={GetRandomSusById(props.product.id)} alt="product-image" className={styles["element-image"]}/>
                <div>
                    <p>{props.product.name}</p>
                    <p>{props.product.price} ₽</p>
                </div>
                <span>x{props.amount}</span>
            </div>
            <Button text="X" onClick={handleButtonClick}/>

        </div>
    )
}

export default ElementView