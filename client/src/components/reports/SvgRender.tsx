interface SvgRenderProps {
    svg: string;
}

export const SvgRender = ({ svg }: SvgRenderProps) => {

    // enconde base64
    let b64_svgtree = btoa(unescape(encodeURIComponent(svg)))

    console.log(svg);

    return (
        <>
            <img src={
                `data:image/svg+xml;base64,${b64_svgtree}`
            } alt="CST" width="100%" height="100%" />
        </>

    );
}