import React from 'react';
import starFull from '../../assets/images/Star Fill.svg';
import starEmpty from '../../assets/images/Star Stroke.svg';

interface IRateValue {
    ratingValue: number,
    maxRatingValue: number
}

const StarRating:React.FC<IRateValue> = ({ratingValue,maxRatingValue}) => {
    const stars = [];
    for(let i = 0; i < maxRatingValue; i++) {
        if(i < ratingValue) {
            stars.push(<img src={starFull}/>);
        } else {
            stars.push(<img src={starEmpty}/>);
        }
    }
    return (
        <>
            {stars}
        </>
    );
};

export default StarRating;
