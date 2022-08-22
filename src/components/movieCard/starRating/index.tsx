import React from 'react';
import Icons from "../../UI/Icons";


interface IRateValue {
    ratingValue: number,
    maxRatingValue: number
}

const StarRating:React.FC<IRateValue> = ({ratingValue,maxRatingValue}) => {
    const stars = [];
    for(let i = 0; i < maxRatingValue; i++) {
        if(i < ratingValue) {
            stars.push(<Icons key={i} name="icon-star-full" width={16} height={16} color="#fff"/>);
        } else {
            stars.push(<Icons key={i} name="icon-star-empty" width={16} height={16} color="#000"/>);
        }
    }
    return (
        <>
            {stars}
        </>
    );
};

export default StarRating;
