import Header from "../header/Header"
import styles from "./MainLayout.module.scss"

type Props = {
    children?: React.ReactNode
}

const MainLayout: React.FC<Props> = ({children}) => {

    return (
        <>
            <Header />
            <div className={styles["content-wrapper"]}>
                {children}
            </div>
        </>
    )
}

export default MainLayout