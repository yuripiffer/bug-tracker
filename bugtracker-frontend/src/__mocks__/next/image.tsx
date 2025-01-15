import React from 'react'
import { ImageProps } from 'next/image'

const NextImage: React.FC<ImageProps> = ({ src, alt, ...props }) => {
    // eslint-disable-next-line @next/next/no-img-element
    return <img src={src as string} alt={alt} {...props} />
}

export default NextImage 