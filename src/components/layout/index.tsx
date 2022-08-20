import React from 'react';
import Header from "../header";
import SideBar from "../sidebar";
import Footer from "../footer";

interface IProps {
    children: React.ReactNode
}

const Layout:React.FC<IProps> = (children:IProps) => {
    return (
        <>
            <Header/>
            <SideBar/>
            {children.children}
            <Footer/>
        </>
    );
};

export default Layout;
