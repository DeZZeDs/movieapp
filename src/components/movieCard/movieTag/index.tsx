import React from 'react';
import styles from './movieTag.module.scss';

interface ITag {
    tagName: string
}
const MovieTag = (tag:ITag) => {
    return (
        <div className={styles.tag}>
            {tag.tagName}
        </div>
    );
};

export default MovieTag;
