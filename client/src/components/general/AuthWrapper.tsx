import React from 'react'
import useCheckSessionQuery from '../../hooks/useCheckSessionQuery'

export type AuthWrapperProps = {
    children: React.ReactNode
}

const AuthWrapper = ({children}: AuthWrapperProps) => {

    const checkSessionQuery = useCheckSessionQuery();

    return (
        <>
            {children}
        </>
    )
}

export default AuthWrapper
