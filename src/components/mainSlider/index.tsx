import React from 'react';
import { Splide, SplideSlide } from '@splidejs/react-splide';
import '@splidejs/react-splide/css';
import slideImg from '../../assets/images/gvk.jpg';
import StarRating from "../starRating";
import MovieTag from '../UI/movieTag';
import styles from './mainSlider.module.scss';


const arrSlides = [{
    tags:['Science Fiction'],
    ratingValue: 5,
    title: 'Godzilla vs. Kong',
    description: 'In a time when monsters walk the Earth, humanity’s fight for its future sets Godzilla and Kong on a collision course that will see the two most powerful forces of nature on the planet collide in a spectacular battle for the ages.',
},
{
    tags:['Science Fiction'],
    ratingValue: 3,
    title: 'Godzilla vs. Kong',
    description: 'In a time when monsters walk the Earth, humanity’s fight for its future sets Godzilla and Kong on a collision course that will see the two most powerful forces of nature on the planet collide in a spectacular battle for the ages.',
}];


const MainSlider:React.FC = () => {

    return (
        <section className={styles.mainSlider}>
            <Splide aria-label="main-slider" options={ {
                rewind: true,
                width : 1920,
                gap   : '1rem',
                perPage: 1,
                breakpoints: {
                    320: {
                        perPage: 1,
                    }
                }
            } }>
                {
                    arrSlides.map((slide, index) => (
                            <SplideSlide key={index} className={styles.mainSlider__slide}>
                                <img src={slideImg} alt="slide"/>
                                <div className={styles.sliderInfo}>
                                    <div className="sliderInfo__wrapper">
                                        <div className="sliderInfo__tags">
                                            {
                                                slide.tags.map((tag, index) => (
                                                    <MovieTag tagName={tag}/>
                                                ))
                                            }
                                        </div>
                                        <div className="sliderInfo__rating">
                                            <StarRating key={index} ratingValue={slide.ratingValue} maxRatingValue={5}/>
                                        </div>
                                        <div className="sliderInfo__title">{slide.title}</div>
                                        <div className="sliderInfo__title">{slide.description}</div>
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
