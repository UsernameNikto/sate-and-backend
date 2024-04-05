import styles from "./page.module.css";
import Menu from "@/components/Menu/Menu";

export const metadata = {
  title: 'Кабинет',
}

export default function Kabinet() {
  return (
    <main className={styles.main}>
      <Menu />

      <div className={styles.center}>
        Кабинет
      </div>
      
      <form>
        <input></input>
        <input></input>
        <button>Войти</button>
      </form>

      </main>
  );
}
