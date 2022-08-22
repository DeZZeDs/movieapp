import React from 'react';
import logo from '../../assets/images/logo.svg';
import Icons from "../UI/Icons";
import styles from './sidebar.module.scss';
import {Link} from "react-router-dom";

const SideBar:React.FC = () => {

    return (
        <div className={styles.sidebar}>
            <div className={styles.sidebar__wrapper}>
                <div className={styles.sidebar__list}>
                    <div className={styles.sidebar__listItem}>
                        <Link to="/">
                            <img className="sidebar__listItemLogo" src={logo} alt="logo"/>
                        </Link>
                    </div>
                    <div className={styles.sidebar__listItem}>
                        <Icons name="icon-search-big" width={32} height={32} color="#fff"/>
                    </div>
                    <div className={styles.sidebar__listItem}>
                        <Link to="/movies/">
                            <Icons name="icon-movie-big" width={32} height={32} color="#fff"/>
                        </Link>

                    </div>
                    <div className={styles.sidebar__listItem}>
                        <Link to="/favourites/">
                            <Icons name="icon-heart-big" width={32} height={32} color="#fff"/>
                        </Link>
                    </div>
                    <div className={styles.sidebar__listItem}>
                        <Link to="/profile:id/">
                            <Icons name="icon-profile-big" width={32} height={32} color="#fff"/>
                        </Link>
                    </div>
                    <div className={styles.sidebar__listItem}>
                        <Link to="/settings/">
                            <Icons name="icon-settings-big" width={32} height={32} color="#fff"/>
                        </Link>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default SideBar;
