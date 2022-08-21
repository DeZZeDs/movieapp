import React from 'react';
import { Splide, SplideSlide } from '@splidejs/react-splide';
import '@splidejs/react-splide/css';
import slideImg from '../../assets/images/gvk.jpg';
import StarRating from "../movieCard/starRating";
import MovieTag from '../movieCard/movieTag';
import styles from './mainSlider.module.scss';
import {Link} from "react-router-dom";


const arrSlides = [{
    tags:['Science Fiction'],
    ratingValue: 5,
    url:'/movies/movie:12124',
    title: 'Godzilla vs. Kong',
    description: 'In a time when monsters walk the Earth, humanity’s fight for its future sets Godzilla and Kong on a collision course that will see the two most powerful forces of nature on the planet collide in a spectacular battle for the ages.',
},
{
    tags:['Science Fiction'],
    ratingValue: 3,
    url:'/movies/movie:12124',
    title: 'Godzilla vs. Kong 2',
    description: 'In a time when monsters walk the Earth, humanity’s fight for its future sets Godzilla and Kong on a collision course that will see the two most powerful forces of nature on the planet collide in a spectacular battle for the ages.',
}];


const MainSlider:React.FC = () => {

    const sliderOptions = {
        rewind: true,
        width: 1920,
        height:800,
        gap: '1 rem',
        perPage: 1,
        breakpoints: {
            320: {
                height: 480,
            },
            576: {
                height: 480
            },
            768: {
                height:600
            },
            1440: {
                height:800
            }
        },
    }
    return (
        <section className={styles.mainSlider}>
            <Splide aria-label="main-slider" options={sliderOptions}>
                {
                    arrSlides.map((slide, index) => (
                            <SplideSlide key={index} className={styles.mainSlider__slide}>
                                <img src={slideImg} alt="slide"/>
                                <div className={styles.overlay}></div>
                                <div className={styles.sliderInfo}>
                                    <div className={styles.sliderInfo__wrapper}>
                                        <div className={styles.sliderInfo__tags}>
                                            {
                                                slide.tags.map((tag, index) => (
                                                    <MovieTag key={index} tagName={tag}/>
                                                ))
                                            }
                                        </div>
                                        <div className={styles.sliderInfo__rating}>
                                            <StarRating key={index} ratingValue={slide.ratingValue} maxRatingValue={5}/>
                                        </div>
                                        <div className={styles.sliderInfo__title}>{slide.title}</div>
                                        <div className={styles.sliderInfo__description}>{slide.description}</div>
                                        <Link className={styles.sliderInfo__link} to={slide.url}>Смотреть</Link>
                                    </div>
                                </div>
                            </SplideSlide>
                        )
                    )
                }
            </Splide>
        </section>
    );
};

export default MainSlider;
