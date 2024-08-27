import styles from "./StoreView.module.scss";
import { useNavigate } from "react-router-dom";
import Button from "../../components/button/Button";
import { GetRandomSusById } from "../../randomSusAssert";
import { CartStoreAndOwnProps, CartStoreProps, cartStoreState, cartStoreStateAndProps, StateProps } from "../../common/cart";
import CartDispatcher, { DispatcherProps } from "../../common/cart/dispatcher";
import { connect } from "react-redux";


const ElementView: React.FC<CartStoreAndOwnProps> = (props) => {    
    const navigate = useNavigate();

    const handleElementClick = () => {
        navigate(`/product/${props.own.id}`)
    }

    const handleButtonClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.stopPropagation()
        props.addToCart(props.own.id)
    }

    return (
        <div className={styles["element-view"]} onClick={handleElementClick}>
            <img src={GetRandomSusById(props.own.id)} alt="product-image" className={styles["element-image"]}/>
            <div>
                <p>{props.own.name}</p>
                <p>{props.own.price} ₽</p>
            </div>
            <Button text="Купить" onClick={handleButtonClick}/>

        </div>
    )
}

export default connect<StateProps, DispatcherProps>(
    cartStoreStateAndProps,
    CartDispatcher,
)(ElementView)