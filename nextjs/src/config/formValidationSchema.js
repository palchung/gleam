import { yupResolver } from '@hookform/resolvers/yup'
import * as Yup from 'yup'

const signupSchema = Yup.object().shape({
    firstName: Yup.string()
        .required('First name is requiredsssssss'),
    lastName: Yup.string()
        .required('Last name is required'),
    email: Yup.string()
        .required('Email is required')
        .email('Email is invalid'),
    password: Yup.string()
        .min(6, 'Password must be at least 6 characters')
        .required('Password is required'),
    confirmPassword: Yup.string()
        .oneOf([Yup.ref('password'), null], 'Passwords must match')
        .required('Confirm Password is required'),
    acceptTerms: Yup.bool()
        .oneOf([true], 'Accept our terms & conditions is required')
})

const loginSchema = Yup.object().shape({
    email: Yup.string()
        .required('Email is required')
        .email('Email is invalid'),
    password: Yup.string()
        .min(6, 'Password must be at least 6 characters')
        .required('Password is required')
})

const signupFormOptions = { resolver: yupResolver(signupSchema) }
const loginFormOptions = { resolver: yupResolver(loginSchema) }

export { signupFormOptions, loginFormOptions }
