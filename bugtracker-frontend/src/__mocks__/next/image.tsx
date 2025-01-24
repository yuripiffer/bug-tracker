import React, { CSSProperties } from 'react'

// const NextImage: React.FC<ImageProps> = ({ src, alt, ...props }) => {
//     // eslint-disable-next-line @next/next/no-img-element
//     return <img src={src as string} alt={alt} {...props} />
// }

// export default NextImage 


jest.mock('next/image', () => ({
    __esModule: true,
    default: ({ src, alt, fill, ...props }: { src: string, alt: string, fill?: boolean }) => {
      const style: CSSProperties = fill ? { position: 'absolute', top: 0, left: 0, bottom: 0, right: 0 } : {};
      // eslint-disable-next-line @next/next/no-img-element
      return <img src={src} alt={alt} style={style} {...props} />;
    }
  }));