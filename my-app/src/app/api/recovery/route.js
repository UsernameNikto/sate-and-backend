import { NextRequest, NextResponse } from "next/server";
export async function POST(request) {
  const { email, newPassword } = await request.json();
  if (!email || !newPassword) {
    return NextResponse.json({ success: false, message: "Пожалуйста, заполните все поля." }, { status: 400 });
  }
  function isValidPassword(password) {
    return password.length >= 6 && /[!@#$%^&*(),.?":{}|<>]/.test(password);
  }
  if (!isValidPassword(newPassword)) {
    return NextResponse.json({ success: false, message: "Пароль должен содержать не менее 6 символов и хотя бы один специальный символ." }, { status: 400 });
  }
  return NextResponse.json({ success: true });
}