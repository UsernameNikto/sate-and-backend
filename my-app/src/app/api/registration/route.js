import { NextRequest, NextResponse } from "next/server";
export async function POST(request) {
  try {
    const { login, password, email, confirmedPassword } = await request.json();
    // Проверка на пустые значения
    if (!login || !password || !email || !confirmedPassword) {
      return NextResponse.json({ success: false, error: "Пожалуйста, заполните все поля" }, { status: 400 });
    }
    // Проверка на совпадение паролей
    if (password !== confirmedPassword) {
      return NextResponse.json({ success: false, error: "Пароли не совпадают" }, { status: 400 });
    }
    // Проверка на длину локальной части email
    const localPart = email.split('@')[0]; // Получаем локальную часть email
    if (localPart.length !== 4) {
      return NextResponse.json({ success: false, error: "Неверная длина локальной части email" }, { status: 400 });
    }
    // Если все проверки пройдены успешно, можно выполнить регистрацию пользователя
    // Здесь можно добавить вашу логику для регистрации пользователя
    return NextResponse.json({ success: true });
  } catch (error) {
    console.error("Error during registration:", error);
    return NextResponse.json({ success: false, error: "Ошибка при регистрации" }, { status: 500 });
  }
}