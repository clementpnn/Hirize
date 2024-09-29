import { createFileRoute, useNavigate } from '@tanstack/react-router'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { Card, CardContent } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { useMutation } from '@tanstack/react-query'
import { formSchema } from '@/types/forms'
import { UserService } from '@/services/user'

export const Route = createFileRoute('/signin')({
  component: Signin,
})

export default function Signin() {
  const navigate = useNavigate({ from: '/signin' })

  const { mutate: signinMutate, isPending: isSignInPending } = useMutation({
    mutationFn: UserService.login,
    onSuccess: () => navigate({ to: '/dashboard' }),
    // onSuccess: (response) => alert(`${response?.message}`),
    onError: (response) => alert(`${response?.message}`),
  })

  const { mutate: signupMutate, isPending: isSignUpPending } = useMutation({
    mutationFn: UserService.create,
    onSuccess: () => navigate({ to: '/dashboard' }),
    // onSuccess: (response) => alert(`${response?.message}`),
    onError: (error) => alert(error || 'Unknown error occurred'),
  })

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: '',
      password: '',
    },
  })

  return (
    <main className="flex flex-col items-center justify-center min-h-screen">
      <Tabs defaultValue="signin" className="w-[400px]">
        <TabsList className="grid w-full grid-cols-2">
          <TabsTrigger value="signin">Se connecter</TabsTrigger>
          <TabsTrigger value="signup">S"inscrire</TabsTrigger>
        </TabsList>
        <TabsContent value="signin">
          <Card>
            <CardContent className="space-y-2 pt-6">
              <Form {...form}>
                <form
                  onSubmit={form.handleSubmit(
                    (values: z.infer<typeof formSchema>) =>
                      signinMutate(values),
                  )}
                  className="space-y-8"
                >
                  <FormField
                    control={form.control}
                    name="email"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Email</FormLabel>
                        <FormControl>
                          <Input
                            placeholder="Entrez votre email"
                            type="email"
                            {...field}
                          />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={form.control}
                    name="password"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Mot de passe</FormLabel>
                        <FormControl>
                          <Input
                            placeholder="Entrez votre mot de passe"
                            type="password"
                            {...field}
                          />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <Button
                    type="submit"
                    disabled={isSignUpPending || isSignInPending}
                  >
                    {isSignUpPending || isSignInPending
                      ? '...'
                      : 'Se connecter'}
                  </Button>
                </form>
              </Form>
            </CardContent>
          </Card>
        </TabsContent>
        <TabsContent value="signup">
          <Card>
            <CardContent className="space-y-2 pt-6">
              <Form {...form}>
                <form
                  onSubmit={form.handleSubmit(
                    (values: z.infer<typeof formSchema>) =>
                      signupMutate(values),
                  )}
                  className="space-y-8"
                >
                  <FormField
                    control={form.control}
                    name="email"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Email</FormLabel>
                        <FormControl>
                          <Input
                            placeholder="Entrez votre email"
                            type="email"
                            {...field}
                          />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={form.control}
                    name="password"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Mot de passe</FormLabel>
                        <FormControl>
                          <Input
                            placeholder="Entrez votre mot de passe"
                            type="password"
                            {...field}
                          />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <Button
                    type="submit"
                    disabled={isSignUpPending || isSignInPending}
                  >
                    {isSignUpPending || isSignInPending ? '...' : "S'inscrire"}
                  </Button>
                </form>
              </Form>
            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>
    </main>
  )
}
