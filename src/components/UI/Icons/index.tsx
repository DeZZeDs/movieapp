import React from 'react';
import IconInterface from './index.d';

const Icons:React.FC<IconInterface> = ({ name, color, width, height }) => {
    return (
        <svg className={`icon ${name}`} width={width} height={height} fill={color}>
            <use xlinkHref={`sprites.svg#${name}`}/>
        </svg>
    );
};

export default Icons;
