import styles from "./page.module.css";
import Menu from "@/components/Menu/Menu";

export const metadata = {
  title: 'Контакты',
}

export default function Contacts() {
  return (
    <main className={styles.main}>
      <Menu />

      <div className={styles.center}>
        Контакты
      </div>
    </main>
  );
}
