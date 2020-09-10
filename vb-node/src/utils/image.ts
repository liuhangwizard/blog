import React from 'react';

export const disableImageDrag=(event: React.DragEvent<HTMLImageElement>):void=>{
    event.preventDefault();
}