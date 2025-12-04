import {
    NavigationMenu,
    NavigationMenuList,
    NavigationMenuItem,
    NavigationMenuLink,
} from '@/components/ui/navigation-menu';

export const GlobalNavigation = () => {
    return (
        <header className="sticky top-0 z-40 w-full border-b bg-background ">
            <div className="container mx-auto flex h-16 items-center space-x-4 sm:justify-between sm:space-x-0">
                <div className="font-black text-foreground text-lg">PUSTAK</div>
                
            
            <NavigationMenu>
                <NavigationMenuList>
                    <NavigationMenuItem>
                        <NavigationMenuLink href="/">Home</NavigationMenuLink>
                    </NavigationMenuItem>
                    <NavigationMenuItem>
                        <NavigationMenuLink href="/about">About</NavigationMenuLink>
                    </NavigationMenuItem>
                    {/* More navigation items */}
                </NavigationMenuList>
            </NavigationMenu>
            </div>
        </header>
    );
}