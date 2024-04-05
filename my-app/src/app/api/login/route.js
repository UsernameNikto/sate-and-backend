import { NextRequest, NextResponse } from "next/server";
export async function POST(request) {
    const {login, password} = await request.json();
    console.log(login, password)
    if (login === "admin" && password === "qwerty123") {
        return NextResponse.json({isLogin:true})
    } else {
        return NextResponse.json({isLogin:false}, {
            status: 400
        })
    }
}