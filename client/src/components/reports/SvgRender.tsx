interface SvgRenderProps {
    svg: string;
}

export const SvgRender = ({ svg }: SvgRenderProps) => {


    return (
        <svg dangerouslySetInnerHTML={{ __html: svg }} className='w-full' />
    );
}