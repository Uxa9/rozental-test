import styles from "./Button.module.scss"


interface ButtonProps extends React.ComponentPropsWithoutRef<"button"> {
    text: string
}

const Button: React.FC<ButtonProps> = (props) => {

    return (
        <button className={styles["button"]} {...props}>
            {props.text}
        </button>
    )
}

export default Button